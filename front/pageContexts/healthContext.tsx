import { createContext } from 'react'
import { UnaryResponse } from '@protobuf/sandbox_pb'
import { CheckServicesResponse } from '@protobuf/health_pb'

export type ContextDataType = {
    result: CheckServicesResponse.AsObject | null,
    dateTime: string,
    revalidate: number,
}

const HealthContext = createContext<ContextDataType>({
    result: null,
    dateTime: '',
    revalidate: 0,
})

export default HealthContext
