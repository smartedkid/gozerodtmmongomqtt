package consumer

import (
	"2112a-6/month/wsyx/cobra/models"
	"2112a-6/month/wsyx/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func Consumer() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("order"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"10.2.171.43:9876"})),
	)
	err := c.Subscribe("order-stock", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

		var order model.Order

		for _, msg := range msgs {
			fmt.Printf("subscribe callback: %v \n", string(msg.Body))
			json.Unmarshal(msg.Body, &order)
			fmt.Println(">>>>>>>>>>>>>>>", order.Goodsid, order.GoodsCount)
			err := models.Db.Exec("update  `goods` set goods_num=goods_num-? where id=?", order.GoodsCount, order.Goodsid).Error
			if err != nil {
				panic(err)
			}
		}

		//todo gorm 连表 插入到es
		//todo 删除key

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	select {}
}
