<template>
    <div id="home">
        <div class="m-8 px-4 md:mx-32 h-full">
            <div class="flex justify-center mt-3 p-5 gap-x-2">
                <input type="text" placeholder="room name" v-model="roomName" class="input-field p-3 rounded-md">
                <button class="py-3 px-4 text-white bg-blue-600 rounded-md font-bold text-lg"
                    @click.prevent="createRoom()">Create room</button>
            </div>
        </div>
        <div class="content">
            <div class="text-left">Available Rooms</div>
            <div class="room_items mt-4 flex flex-wrap gap-3">
                <div v-for="item of rooms" :key="item.id" class="room_item">
                    <div class="mr-2">
                        <div class="text-left text-gray-500 text-sm">Room</div>
                        <div class="text-left mt-2 text-lg text-ellipsis">{{ item.name }}</div>
                    </div>
                    <div class="flex items-center">
                        <button class="py-2 px-4 text-white bg-blue-600 rounded-md text-base"
                            @click.prevent="joinRoom(item.id)">Join</button>
                    </div>
                </div>

            </div>
        </div>
    </div>
</template>

<script>
// @ is an alias to /src

export default {
    name: 'HomeView',
    components: {
    },
    data() {
        return {
            rooms: [],
            roomName: ''
        }
    },
    async mounted() {
        try {
            let res = await fetch('http://localhost:8000/ws/getRooms', {
                method: "GET",
                headers: { 'Content-Type': 'application/json' },
            })

            let data = await res.json()
            if (res.ok) {
                this.rooms = data
            } else {
                console.log(data.error)
            }
        } catch (e) {
            console.log(e)
        }
    },
    methods: {
        async createRoom() {
            if (this.roomName == '') {
                console.log('room name is empty')
                return
            }
            let room = {
                id: String(this.rooms.length + 1),
                name: this.roomName
            }
            let res = await fetch('http://localhost:8000/ws/createRoom', {
                method: "POST",
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(room)
            })

            let data = await res.json()
            if (res.ok) {
                this.rooms.push(data)
            }
        },
        joinRoom(roomId) {
            console.log(roomId)
            this.$router.push(`/app/${roomId}`)
        }
    }
}
</script>

<style scoped>
.input-field {
    border: 2px solid rgb(169, 169, 169, .6);
    outline: none;
}

.input-field:focus {
    border: 2px solid rgb(54, 98, 227);
}

.content {
    margin: 5rem 10rem;
    font-weight: 600;
    font-size: 16px;
}

.room_items {}

.room_item {
    padding: 10px 15px;
    width: 350px;
    height: 90px;
    border: 2px solid rgb(54, 98, 227, .6);
    border-radius: 10px;
    display: flex;
    justify-content: space-between;
}
</style>
