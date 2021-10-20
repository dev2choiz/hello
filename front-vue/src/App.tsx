import { defineComponent } from 'vue'

export default defineComponent({
    setup () {
        return () => <div id="nav">
            <router-link to="/">Home</router-link>
            <router-link to="/about">About</router-link>
            <router-view/>
        </div>
    }
})
