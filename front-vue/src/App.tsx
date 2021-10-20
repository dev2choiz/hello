import { defineComponent } from 'vue'
import { RouterView, RouterLink } from 'vue-router'

export default defineComponent({
    setup () {
        return () => <div id="nav">
            <RouterLink to="/">Unary</RouterLink><br/>
            <RouterLink to="/server-stream">ServerStream</RouterLink><br/>
            <RouterView/>
        </div>
    }
})
