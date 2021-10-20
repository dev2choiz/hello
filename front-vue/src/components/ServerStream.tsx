import { defineComponent, onBeforeUnmount, onMounted, ref } from 'vue'
import grpcSvc from '@services/grpcSvc'

export default defineComponent({
    name: 'ServerStream',

    setup () {
        let cancel: (() => void) | null = null
        const data = ref<Array<string>>([])

        onMounted(async () => {
            const res = await grpcSvc.sandboxServerStream()
            // receive stream data from the server
            res.on('data', msg => {
                data.value.push(msg.getMessage())
            })
            cancel = res.cancel
        })

        onBeforeUnmount(() => {
            // stop the stream on unmount
            if (cancel) cancel()
        })

        return () => <div>
            { data.value.map(v => <div>{v}</div>) }
        </div>
    }
})
