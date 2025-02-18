import type { GetServerSidePropsContext, InferGetServerSidePropsType, NextPage } from 'next'
import getConfig from 'next/config'
import UnaryComp from '@components/UnaryComp'
import { UnaryRequest, UnaryResponse } from '@protobuf/sandbox_pb'
import { SandboxClient } from '@protobuf/sandbox_pb_service'
import { grpc } from '@improbable-eng/grpc-web'
import { NodeHttpTransport } from '@improbable-eng/grpc-web-node-http-transport'
import IndexContext from '@/pageContexts/indexContext'
import headerSvc from '@/services/headerSvc'

const UnaryPage: NextPage<InferGetServerSidePropsType<typeof getServerSideProps>> = (props) => {
    return <IndexContext.Provider value={props}>
        <UnaryComp />
    </IndexContext.Provider>
}

export async function getServerSideProps(ctx: GetServerSidePropsContext) {
    const req = new UnaryRequest()
    const name = (ctx.query.name ?? '') as string
    req.setName(name)
    const opts = {} as grpc.RpcOptions
    if ('undefined' === typeof window) {
        opts.transport = NodeHttpTransport()
        opts.debug = false
    }
    const { serverRuntimeConfig } = getConfig()

    const client = new SandboxClient(serverRuntimeConfig.serverGrpcBaseUrl, opts)
    const result = await new Promise<UnaryResponse.AsObject | null>(resolve => {
        client.unary(req, (error, res) => {
            if (!!error) {
                console.error(error.message)
                resolve(null)
                return
            }
            resolve((res as UnaryResponse).toObject())
        })
    })

    ctx.res.setHeader('Cache-Control', headerSvc.getCacheControl(3600))

    return {
        props: {
            result,
        },
    }
}

export default UnaryPage
