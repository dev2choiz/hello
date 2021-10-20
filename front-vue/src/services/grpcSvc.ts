import { grpc } from '@improbable-eng/grpc-web'
import { UnaryRequest, UnaryResponse } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'

class GrpcSvc {
    private baseUrl: string = process.env.VUE_APP_GRPC_BASE_URL

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
}

const grpcSvc = new GrpcSvc()
export default grpcSvc
