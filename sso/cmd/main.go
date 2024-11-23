package main

import (
	"fmt"
	// "github.com/LavaJover/auth-project/auth-service/gen/authpb"
	"github.com/LavaJover/auth-project/sso/internal/config"
)

func main(){
	cfg := config.MustLoad()
	fmt.Println(cfg)
}