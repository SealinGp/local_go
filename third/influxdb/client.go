package influxdb

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	log2 "github.com/influxdata/influxdb-client-go/v2/log"
)

var (
	//org & bucket ? 表示查询时间的一种策略
	org    = "my-org"
	bucket = "mydb"

	//表名
	measurement = "cpu"

	//例子数据
	i    = 0
	tags = map[string]string{
		"host": "localhost",
	}
	fields = map[string]interface{}{
		"region": "china-huadong",
		"value":  0.55,
	}
)

func NewInfluxDbClient(serverURL, authToken string, logger *log.Logger, ctx context.Context) *influxdbClient {
	ifdbClient := &influxdbClient{log: logger, ctx: ctx}
	options := influxdb2.DefaultOptions().
		SetUseGZip(true).
		SetTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}).
		SetBatchSize(4).SetLogLevel(log2.DebugLevel)
	ifdbClient.client = influxdb2.NewClientWithOptions(serverURL, authToken, options)
	return ifdbClient
}

type influxdbClient struct {
	log    *log.Logger
	client influxdb2.Client
	ctx    context.Context
}

//获取例子数据
func (ifdbClient *influxdbClient) GetExampleData() (tagsE map[string]string, fieldsE map[string]interface{}) {
	tagsE = make(map[string]string)
	for tagK, tagV := range tags {
		tagsE[tagK] = tagV + strconv.Itoa(i)
	}
	fieldsE = make(map[string]interface{})

	for fieldK, fieldV := range fields {
		switch v := fieldV.(type) {
		case string:
			fieldsE[fieldK] = v + strconv.Itoa(i)
		case float32:
			fieldsE[fieldK] = v + float32(i)
		case float64:
			fieldsE[fieldK] = v + float64(i)
		}
	}

	i++
	return
}

//写入数据
//point 表示一条数据 包含measurement,a tag set,a field key,a field value,a timestamp属性
func (ifdbClient *influxdbClient) Write(ifBlock bool) {

	//阻塞插入
	if ifBlock {
		ifdbClient.log.Println("write block")
		apiWriter := ifdbClient.client.WriteAPIBlocking(org, bucket)

		//插入一条数据的两种方式
		//方式1: map
		tags, fields := ifdbClient.GetExampleData()
		point := influxdb2.NewPoint(measurement, tags, fields, time.Now())
		err := apiWriter.WritePoint(ifdbClient.ctx, point)
		if err != nil {
			//ifdbClient.log.Println("insert err:",err)
		}

		//方式2:链式操作
		tags, fields = ifdbClient.GetExampleData()
		point = influxdb2.NewPointWithMeasurement(measurement)
		for tagK, tagV := range tags {
			point.AddTag(tagK, tagV)
		}
		for fieldK, fieldV := range fields {
			point.AddField(fieldK, fieldV)
		}
		point.SetTime(time.Now())
		err = apiWriter.WritePoint(ifdbClient.ctx, point)
		if err != nil {
			//ifdbClient.log.Println("insert err:",err)
		}
		return
	}

	//批量非阻塞插入,最大一次插入20条(在初始化的时候设置了)
	ifdbClient.log.Println("write not block")
	apiWriter := ifdbClient.client.WriteAPI(org, bucket)
	for i := 0; i < 5; i++ {
		tags, fields := ifdbClient.GetExampleData()
		point := influxdb2.NewPoint(measurement, tags, fields, time.Now())
		apiWriter.WritePoint(point)
	}
	apiWriter.Flush()

}
func (ifdbClient *influxdbClient) Query() {

	queryApi := ifdbClient.client.QueryAPI(org)
	result, err := queryApi.Query(ifdbClient.ctx, `from(bucket:"`+bucket+`") 
|> range(start: -12h,stop: now()) 
|> filter(fn: (r) => r.measurement == "cpu")`)
	if err != nil {
		ifdbClient.log.Println("query err:", err)
		return
	}
	defer result.Close()

	for result.Next() {
		if result.TableChanged() {
			ifdbClient.log.Println("table:", result.TableMetadata().String())
		}
		fmt.Println("?")
		ifdbClient.log.Println("row:", result.Record().String())
	}
	if result.Err() != nil {
		ifdbClient.log.Println("result err:", result.Err().Error())
	}
}
func (ifdbClient *influxdbClient) Close() {
	ifdbClient.client.Close()
}
