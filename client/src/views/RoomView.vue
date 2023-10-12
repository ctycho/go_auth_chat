<template>
    <div class="app_container">
        <div class="block">
            <div class="title">My Chat</div>
            <div v-if="!messages.length" class="empty">No messages here yet...</div>
            <div id="content" class="content">
                <div class="messages">
                    <div v-for="item of messages" :key="item.username">
                        <MessageItem :item="item" :myID="username" />
                    </div>
                </div>
            </div>
        </div>
        <div class="interact">
            <input type="text" class="input" v-model="text" v-on:keyup.enter="sendMessage()"
                placeholder="Text a message...">
            <button class="button" v-on:click="sendMessage()">Send</button>
        </div>
    </div>
</template>

<script>
import MessageItem from '../components/MessageItem.vue'
import { WEBSOCKET_URL } from '../../constants/index'

console.log(WEBSOCKET_URL)

export default {
    name: 'RoomView',
    components: {
        MessageItem
    },
    data() {
        return {
            socket: null,
            messages: [],
            text: '',
            userID: '',
            username: '',
            roomId: '',
        }
    },
    methods: {
        sendMessage() {
            if (!this.text) return
            console.log(this.text)
            this.socket.send(this.text)
            this.text = ''
        },
        getTime() {
            const d = new Date()
            let hours = d.getHours()
            let min = d.getMinutes()
            return `${hours < 10 ? `0${hours}` : hours}:${min < 10 ? `0${min}` : min}`
        }
    },
    mounted() {
        this.roomId = this.$route.params['roomId']
        try {
            const obj = localStorage.getItem('user_info')
            if (!obj) {
                this.$router.push('/login')
            }
            const user = JSON.parse(obj)
            this.userID = user.id
            this.username = user.username
        } catch (e) {
            console.log(e)
        }
        // console.log(`${WEBSOCKET_URL}/ws/joinRoom/${this.roomId}?userId=${this.userID}&username=${this.username}`)

        this.socket = new WebSocket(`${WEBSOCKET_URL}/ws/joinRoom/${this.roomId}?userId=${this.userID}&username=${this.username}`)

        this.socket.onopen = () => {
            console.log('Successfully')
        }

        this.socket.onmessage = (msg) => {
            console.log('Message: ', msg)

            let data = JSON.parse(msg.data)
            console.log(data)
            this.messages.push({
                client: data.username,
                content: data.content,
                time: this.getTime()
            })
            console.log(JSON.parse(JSON.stringify(this.messages)))
        }

        this.socket.onclose = (event) => {
            console.log('Connection closed: ', event)
        }

        this.socket.onerror = (err) => {
            console.log('Connection error: ', err)
        }
    }
}
</script>

<style scoped>
.app_container {
    /* background-color: #e3e3e3; */
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    grid-row-gap: 30px;
    height: 100vh;
}

.block {
    background-color: #e3e3e3;
    border: 2px solid rgb(169, 169, 169, .6);
    /* background: #fff; */
    padding: 25px;
    border-radius: 10px;
    width: 45%;
    height: 50%;
    overflow: hidden;
}

.title {
    font-size: 26px;
    font-weight: 600;
    text-align: center;
    margin-bottom: 25px;
}

.content {
    height: calc(100% - 35.5px - 25px);
    overflow-y: scroll;
}

.messages {
}

.interact {
    width: calc(40% + 50px);
    display: grid;
    grid-template-columns: 8fr 2fr;
    grid-column-gap: 15px;
}

.input {
    padding: 10px 20px;
    outline: none;
    border: 2px solid rgb(169, 169, 169, .6);
    border-radius: 10px;
}

.button {
    border: 1px solid rgb(54, 98, 227);
    background: rgb(54, 98, 227);
    color: #fff;
    border-radius: 10px;
    transition: .3s ease;
}

.button:hover {
    cursor: pointer;
    margin-top: -3px;
    margin-bottom: 3px;
}
</style>
