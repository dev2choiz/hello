import { defineComponent } from 'vue'
import ServerStream from '@components/ServerStream'

export default defineComponent({
    name: 'ServerStreamPage',
    setup () {
        return () => <div>
            <ServerStream />
        </div>
    }
})
