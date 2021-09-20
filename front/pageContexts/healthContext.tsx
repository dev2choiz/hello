import { createContext } from 'react'
import { UnaryResponse } from '@protobuf/sandbox_pb'
import { CheckServicesResponse } from '@protobuf/health_pb'

export type ContextDataType = {
    result: CheckServicesResponse.AsObject | null,
    dateTime: string,
}

const HealthContext = createContext<ContextDataType>({
    result: null,
    dateTime: '',
})

export default HealthContext
