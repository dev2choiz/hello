import { defineComponent } from 'vue'
import Unary from '@components/Unary'

export default defineComponent({
    name: 'Homepage',
    setup () {
        return () => <div>
            <Unary />
        </div>
    }
})
