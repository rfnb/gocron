package rpc

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/google/logger"
	croninit "github.com/qiusnay/gocron/init"
	"github.com/qiusnay/gocron/model"
	"github.com/qiusnay/gocron/service/notify"
	"github.com/qiusnay/gocron/service/rpc/etcd"
	gocron "github.com/qiusnay/gocron/service/rpc/protofile"
	"github.com/qiusnay/gocron/utils"
	"google.golang.org/grpc"
)

var (
	addr        = flag.String("addr", utils.GetLocalIP()+":8973", "server address")
	etcdAddr    = flag.String("etcdAddr", "127.0.0.1:2379", "etcd address")
	ServiceName = flag.String("ServiceName", "task", "service name")
)

const (
	CronNormal  int64 = 10000 //正常
	CronError   int64 = 10002 // 失败
	CronSucess  int64 = 10001 // 成功
	CronTimeOut int64 = 10003 // 超时
)

type CronResponse struct {
	Host   string
	Code   int64
	Result string
	Err    error
}

func Start() {
	fmt.Println("grpc start")
	flag.Parse()
	//定义rpc服务端
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	} else {
		fmt.Printf("listen at:8973\n")
	}
	defer lis.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	gocron.RegisterTaskServer(s, &server{})
	fmt.Printf("server addr:%s\n", *addr)

	//服务注册，go协程for循环定时往etcd上注册服务信息
	go etcd.Register(*EtcdAddr, *ServiceName, *addr, 5)

	//进程终止信号，注销etcd上的服务
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		etcd.UnRegister(*ServiceName, *addr)

		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}

	}()

	//拉起rpc服务
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %s", err)
	}
}

type server struct{}

func (s *server) Run(ctx context.Context, req *gocron.TaskRequest) (*gocron.TaskResponse, error) {
	// logger.Info(fmt.Sprintf("接收通道 : %+v", ctx))
	queryResult := CronResponse{}
	queryCmd := AssembleCmd(req)
	switch req.Querytype {
	case "wget":
	case "curl":
		rpccurl := RpcServiceCurl{}
		queryResult = rpccurl.ExecCurl(ctx, queryCmd)
		break
	default:
		rpcshell := RpcServiceShell{}
		queryResult = rpcshell.ExecShell(ctx, queryCmd, req.Taskid)
	}
	queryResult.Host = utils.GetLocalIP()
	//写入执行日志
	s.AfterExecJob(queryResult, req)

	logger.Info(fmt.Sprintf("execute cmd end: [cmd: %s err: %s, status : %d]",
		queryCmd,
		queryResult.Err,
		queryResult.Code,
	))
	return &gocron.TaskResponse{Err: queryResult.Err.Error(), Output: queryResult.Result, Status: queryResult.Code, Host: queryResult.Host}, nil
}

func (s *server) AfterExecJob(queryResult CronResponse, req *gocron.TaskRequest) {
	var TaskResult = model.TaskResult{}
	TaskResult.Result = queryResult.Result
	TaskResult.Host = queryResult.Host
	TaskResult.Status = queryResult.Code
	TaskResult.Endtime = time.Now().Format("2006-01-02 15:04:05")
	if queryResult.Err != nil {
		TaskResult.Err = queryResult.Err.Error()
	} else {
		TaskResult.Err = "success"
	}
	_, err := new(model.FlLog).UpdateTaskLog(req.Taskid, TaskResult)
	if err != nil {
		logger.Error("任务结束#更新任务日志失败-", err)
	}
	jobModel := model.FlCron{}
	JobInfo, _ := jobModel.GetJobInfo(req.GetJobid())

	// 发送邮件
	go SendNotification(JobInfo[0], TaskResult)
}

// 发送任务结果通知
func SendNotification(jobModel model.FlCron, taskResult model.TaskResult) {
	if taskResult.Err == "succss" {
		return // 执行失败才发送通知
	}
	//发送邮件
	notify.SendCronAlarmMail(taskResult, jobModel)
}

func AssembleCmd(cron *gocron.TaskRequest) string {
	LogFile := GetLogFile(cron.Jobid, cron.Taskid)
	// if utils.IsFile(LogFile) {
	// 	s, err := os.Stat(LogFile)
	// 	s.Chmod(0664)
	// }
	return cron.Command + " > " + LogFile
}

func GetLogFile(Jobid int64, Taskid int64) string {
	//设置日志目录
	LogDir := croninit.BASEPATH + "/log/cronlog/" + time.Now().Format("2006-01-02")
	if !utils.IsDir(LogDir) {
		// mkdir($LogDir, 0777, true);
		os.MkdirAll(LogDir, os.ModePerm)
	}
	StrJobid := strconv.FormatInt(Jobid, 10)
	StrTaskid := strconv.FormatInt(Taskid, 10)
	return LogDir + "/cron-task-" + StrJobid + "-" + StrTaskid + "-log.log"
}
