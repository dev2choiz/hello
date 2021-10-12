import { createContext } from 'react'
import { UnaryResponse } from '@protobuf/sandbox_pb'

export type ContextDataType = {
    result: UnaryResponse.AsObject | null,
    dateTime: string,
    revalidate: number,
}

const UnaryStaticContext = createContext<ContextDataType>({
    result: null,
    dateTime: '',
    revalidate: 0,
})

export default UnaryStaticContext
