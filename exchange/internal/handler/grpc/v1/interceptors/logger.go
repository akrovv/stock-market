package interceptors

import (
	"context"

	"github.com/akrovv/exchange/pkg/logger"
	"google.golang.org/grpc"
)

func Logger(logger logger.Logger) grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req any,
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {

		logger.Infof("[Server Interceptor] method=%s", info.FullMethod)
		m, err := handler(ctx, req)
		logger.Infof("[Server Interceptor] send_data=%v, error=%w", m, err)

		return m, err
	})
}
