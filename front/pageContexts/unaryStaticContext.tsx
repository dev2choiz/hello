import { createContext } from 'react'
import { UnaryResponse } from '@protobuf/sandbox_pb'

export type ContextDataType = {
    result: UnaryResponse.AsObject | null,
    dateTime: string,
}

const UnaryStaticContext = createContext<ContextDataType>({
    result: null,
    dateTime: '',
})

export default UnaryStaticContext
