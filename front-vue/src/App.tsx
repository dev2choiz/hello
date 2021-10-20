import { defineComponent } from 'vue'

export default defineComponent({
    setup () {
        return () => <div id="nav">
            <router-link to="/">Unary</router-link><br/>
            <router-link to="/server-stream">ServerStream</router-link><br/>
            <router-view/>
        </div>
    }
})
