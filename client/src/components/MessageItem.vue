<template>
    <div class="message" :class="{ my_message: myID == item.client }">
        <div class="client">{{ item.client }}</div>
        <div class="flex gap-3 items-baseline">
            <div class="text">{{ item.content }}</div>
            <div class="time">{{ item.time }}</div>
        </div>
    </div>
</template>

<script>

export default {
    name: 'MessageItem',
    props: ['item', 'myID']
}
</script>

<style scoped>
.message {
    padding: 5px 8px;
    border-radius: 5px;
    background-color: #fff;
    max-width: 60%;
    width: max-content;
    grid-column-gap: 7px;
    margin-bottom: 10px;
    line-height: 20px;
}

.my_message {
    margin-left: auto;
}

.message:not(:last-child) {
    /* Doesn't work */
    margin-bottom: 10px;
}

.client {
    text-align: left;
    color: rgb(124, 71, 207);
    font-size: 16px;
}

/* .client span {
    background: #fff;
    border-radius: 50%;
    width: 20px;
    height: 20px;
    font-size: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
} */

.text {
    text-align: left;
    font-size: 15px;
}

.time {
    display: flex;
    align-items: end;
    font-size: 10px;
}
</style>