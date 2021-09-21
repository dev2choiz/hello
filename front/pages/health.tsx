import type { NextPage, InferGetStaticPropsType, GetStaticProps } from 'next'
import HealthComp from '@components/HealthComp'
import config from '@config/config'
import { grpc } from '@improbable-eng/grpc-web'
import { NodeHttpTransport } from '@improbable-eng/grpc-web-node-http-transport'
import HealthContext, { ContextDataType } from '@/pageContexts/healthContext'
import { CheckServicesRequest, CheckServicesResponse } from '@protobuf/health_pb'
import { HealthClient } from '@protobuf/health_pb_service'
import getConfig from 'next/config'

const HealthPage: NextPage<InferGetStaticPropsType<typeof getStaticProps>> = props => {
    return <HealthContext.Provider value={ props }>
        <HealthComp />
    </HealthContext.Provider>
}

export const getStaticProps: GetStaticProps<ContextDataType> = async() => {
    console.log('\ngenerate status\n')

    const opts: grpc.RpcOptions = {}
    opts.transport = NodeHttpTransport()

    const { serverRuntimeConfig } = getConfig()

    const md = new grpc.Metadata()
    md.set('x-api-key', serverRuntimeConfig.apiKey)
    const req = new CheckServicesRequest()
    const client = new HealthClient(config.serverGrpcBaseUrl, opts)
    const result = await new Promise<CheckServicesResponse.AsObject | null>(resolve => {
        client.check(req, md, (error, res) => {
            if (!!error) {
                console.log('error', error)
                resolve(null)
                return
            }
            resolve((res as CheckServicesResponse).toObject())
        })
    })

    return {
        props: { result, dateTime: new Date().toLocaleString() },
        revalidate: 30,
    }
}

export default HealthPage
