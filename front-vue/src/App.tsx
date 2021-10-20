import { defineComponent } from 'vue'

export default defineComponent({
    setup () {
        return () => <div id="nav">
            <router-link to="/">Unary</router-link>
            <router-view/>
        </div>
    }
})
