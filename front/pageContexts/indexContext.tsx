import { createContext } from 'react'
import { UnaryResponse } from '@protobuf/sandbox_pb'

export type ContextDataType = {
    result: UnaryResponse.AsObject | null,
}

const IndexContext = createContext<ContextDataType>({
    result: null,
})

export default IndexContext
