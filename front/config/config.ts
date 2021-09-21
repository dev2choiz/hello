export default {
    // todo: use publicRuntimeConfig
    grpcBaseUrl: process.env.NEXT_PUBLIC_API_BASE_URL as string,
    // todo: use serverRuntimeConfig
    serverGrpcBaseUrl: process.env.SERVER_API_BASE_URL as string,
    unaryStaticParams: ['', 'rand', 'richard', 'fitz', 'belgarion', 'gerald',],
}
