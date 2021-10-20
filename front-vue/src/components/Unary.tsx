import { defineComponent, onMounted, ref } from 'vue'
import grpcSvc from '@services/grpcSvc'
import { UnaryResponse } from '@protobuf/sandbox_pb'
import { useRoute } from 'vue-router'

export default defineComponent({
    name: 'Unary',

    setup () {
        const data = ref<UnaryResponse.AsObject | null>(null)
        const route = useRoute()
        onMounted(async () => {
            data.value = await grpcSvc.sandboxUnary(route.query.name as string || 'World')
        })

        return () => <div>
            <pre>{ JSON.stringify(data.value, null, 2) }</pre>
        </div>
    }
})
