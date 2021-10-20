import { defineComponent, onBeforeUnmount, onMounted, ref } from 'vue'
import grpcSvc from '@services/grpcSvc'

export default defineComponent({
    name: 'ServerStream',

    setup () {
        let cancel: (() => void) | null = null
        const data = ref<Array<string>>([])
        onMounted(async () => {
            const res = await grpcSvc.sandboxServerStream()
            res.on('data', msg => {
                data.value.push(msg.getMessage())
            })
            cancel = res.cancel
        })

        onBeforeUnmount(() => {
            if (cancel) cancel()
        })

        return () => <div>
            <pre>{ JSON.stringify(data.value, null, 2) }</pre>
        </div>
    }
})
