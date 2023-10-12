<template>
    <div class="flex items-center justify-center min-w-full min-h-screen">
        <form action="" class="flex flex-col md:w-1/5">
            <div class="text-3xl font-bold text-center">
                <span class="text-blue-600">welcome!</span>
            </div>
            <input type="text" placeholder="username" v-model="username" class="input-field p-3 mt-8 rounded-md" />
            <input type="text" placeholder="email" v-model="email" class="input-field p-3 mt-4 rounded-md" />
            <input type="password" placeholder="password" v-model="password" class="input-field p-3 mt-4 rounded-md" />
            <a class="mt-2 text-left text-blue-600 underline hover:cursor-pointer">
                <router-link to="/login">Already have an account?</router-link>
            </a>
            <my-button name="sign" @click.prevent="signupUser()"></my-button>
        </form>
    </div>
</template>
  
<script>
// @ is an alias to /src
import MyButton from '@/components/MyButton.vue'
import { API_URL } from '../../constants/index'

export default {
    name: '',
    components: {
        MyButton
    },
    data() {
        return {
            email: '',
            password: '',
            username: ''
        }
    },
    methods: {
        async signupUser() {
            console.log(this.username)
            console.log(this.email)
            console.log(this.password)
            try {
                let res = await fetch(`${API_URL}/signup`, {
                    method: "POST",
                    headers: { 'Content-Type' : 'application/json'},
                    body: JSON.stringify({
                        'username': this.username, 
                        'email': this.email, 
                        'password': this.password
                    })
                })

                let data = await res.json()
                console.log(data)
                console.log(res.ok)
                if (res.ok) {
                    const user = {
                        id: data.id,
                        username: data.username,
                    }

                    this.$router.push('/login')
                } else {
                    console.log(data.error)
                }
            } catch (e) {
                console.log(e)
            }

            this.username = ''
            this.email = ''
            this.password = ''
        }

    }
}
</script>

<style scoped>
.input-field {
    border: 2px solid rgb(169,169,169, .6);
    outline: none;
}
.input-field:focus {
    border: 2px solid rgb(54,98,227);
}
</style>
