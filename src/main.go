package main

import (
	"fmt"
	"net/http"
	"os"
	"go.uber.org/zap"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	cfg "github.com/xiphiar/secret-validator_exporter/config"
	"github.com/xiphiar/secret-validator_exporter/exporter"
	rpc "github.com/xiphiar/secret-validator_exporter/getData/rpc"
)

var (
	bech32MainPrefix = os.Getenv("PREFIX")
)

func main() {

	port := "26661"

	log,_ := zap.NewDevelopment()
        defer log.Sync()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(bech32MainPrefix, bech32MainPrefix+sdk.PrefixPublic)
	config.SetBech32PrefixForValidator(bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator, bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic)
	config.SetBech32PrefixForConsensusNode(bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus, bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic)

	cfg.ConfigPath = "/abc/" //os.Args[1]

	cfg.Init()
	rpc.OpenSocket(log)

	http.Handle("/metrics", promhttp.Handler())
	go exporter.Start(log)

	err := http.ListenAndServe(":" +port, nil)
	// log
        if err != nil {
                // handle error
                log.Fatal("HTTP Handle", zap.Bool("Success", false), zap.String("err", fmt.Sprint(err),))
        } else {
		log.Info("HTTP Handle", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Listen&Serve", "Prometheus Handler(Port: " +port +")"),)
        }

//	utils.Log(err, "Http Handle", "", "")
}
