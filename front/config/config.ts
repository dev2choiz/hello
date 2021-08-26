export default {
    grpcBaseUrl: process.env.NEXT_PUBLIC_API_BASE_URL as string,
    serverGrpcBaseUrl: process.env.SERVER_API_BASE_URL as string,
    unaryStaticParams: ['', 'rand', 'richard', 'fitz', 'belgarion', 'gerald',],
}
