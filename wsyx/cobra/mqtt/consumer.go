package mqtt

import (
	"2112a-6/month/wsyx/cobra/global"
	"2112a-6/month/wsyx/cobra/models"
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
)

var Message mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	//数据过滤
	var data bson.M
	err := json.Unmarshal(msg.Payload(), &data)
	if err != nil {
		panic(err)
	}
	//attribute := data["data"].(map[string]interface{})["attribute"].(string)
	//fmt.Println(attribute)
	fmt.Println(data["data"].(map[string]interface{})["arguments"].(map[string]interface{})["attribute"])

	i := data["value"]
	collection := models.Client.Database("wsyx").Collection("user")
	//_, err = collection.InsertOne(context.Background(), data)
	//if err != nil {
	//	return
	//}

	//_, err = collection.DeleteOne(context.Background(), data)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = collection.UpdateMany(context.Background(), data, bson.M{
	//	"$set": bson.M{
	//		"value": "1111",
	//	},
	//})
	fmt.Println(">>>>>>>>>")
	var result bson.M
	collection.FindOne(context.Background(), bson.M{
		"value": i,
	}).Decode(&result)
	if err != nil {
		return
	}

	fmt.Println(result)

}

func ConsumerMqtt() {

	opts := mqtt.NewClientOptions().AddBroker(global.AddBroker).SetClientID(global.SetClientID).SetUsername(global.User).SetPassword(global.Password)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(Message)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	// 订阅主题
	if token := c.Subscribe(global.Subscribe, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	defer func() {

		// 取消订阅
		if token := c.Unsubscribe(global.Subscribe); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}

		// 断开连接
		c.Disconnect(250)
	}()
	select {}
}
