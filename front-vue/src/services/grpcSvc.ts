import { grpc } from '@improbable-eng/grpc-web'
import { ServerStreamRequest, ServerStreamResponse, UnaryRequest, UnaryResponse } from '@protobuf/sandbox_pb'
import { ResponseStream, SandboxClient } from '@protobuf/sandbox_pb_service'

class GrpcSvc {
    private baseUrl: string = process.env.VUE_APP_GRPC_BASE_URL

    /**
     * Request the server
     * @param name
     */
    public sandboxUnary = async (name: string): Promise<UnaryResponse.AsObject | null> => {
        const req = new UnaryRequest()
        req.setName(name)
        const opts: grpc.RpcOptions = { debug: false }

        const client = new SandboxClient(this.baseUrl, opts)
        return new Promise<UnaryResponse.AsObject | null>(resolve => {
            client.unary(req, (error, res) => {
                if (error) {
                    console.error(error.message)
                    resolve(null)
                    return
                }
                resolve((res as UnaryResponse).toObject())
            })
        })
    }

    /**
     * Do a stream call to the gRPC server
     */
    public sandboxServerStream = (): ResponseStream<ServerStreamResponse> => {
        const req = new ServerStreamRequest()
        req.setNumber(10)
        req.setMsPerResponse(1000)

        const client = new SandboxClient(this.baseUrl)
        return client.serverStream(req)
    }
}

const grpcSvc = new GrpcSvc()
export default grpcSvc
