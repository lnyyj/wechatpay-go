// Copyright 2021 Tencent Inc. All rights reserved.
//
// 服务商批量转账API
//
// No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
// API version: 0.0.2

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package partnertransferbatch_test

import (
	"context"
	"log"

	"github.com/lnyyj/wechatpay-go/core"
	"github.com/lnyyj/wechatpay-go/core/option"
	"github.com/lnyyj/wechatpay-go/services/partnertransferbatch"
	"github.com/lnyyj/wechatpay-go/utils"
)

func ExampleTransferDetailApiService_GetTransferDetailByNo() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := partnertransferbatch.TransferDetailApiService{Client: client}
	resp, result, err := svc.GetTransferDetailByNo(ctx,
		partnertransferbatch.GetTransferDetailByNoRequest{
			BatchId:  core.String("1030000071100999991182020050700019480001"),
			DetailId: core.String("1040000071100999991182020050700019500100"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetTransferDetailByNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleTransferDetailApiService_GetTransferDetailByOutNo() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Print("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}

	svc := partnertransferbatch.TransferDetailApiService{Client: client}
	resp, result, err := svc.GetTransferDetailByOutNo(ctx,
		partnertransferbatch.GetTransferDetailByOutNoRequest{
			OutBatchNo:  core.String("plfk2020042013"),
			OutDetailNo: core.String("x23zy545Bd5436"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetTransferDetailByOutNo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
