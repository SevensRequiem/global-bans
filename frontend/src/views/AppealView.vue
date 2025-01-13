<template>
    <div class="appeal">
        <h2>Appeal</h2>
        <form @submit.prevent="submitAppeal">
        <div class="field">
            <label for="username">Username</label>
            <input type="text" id="username" v-model="username" required />
        </div>
        <div class="field">
            <label for="steamid">Steam ID</label>
            <input type="text" id="steamid" v-model="steamid" required />
        </div>
        <div class="field">
            <label for="minecraftuuid">Minecraft UUID</label>
            <input type="text" id="minecraftuuid" v-model="minecraftuuid" required />
        </div>
        <div class="field">
            <label for="reason">Why should you be unbanned?</label>
            <textarea id="reason" v-model="reason" required></textarea>
        </div>
        <button type="submit">Submit</button>
        </form>
    </div>
</template>


<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from 'axios'

export default defineComponent({
    setup() {
        const username = ref('')
        const reason = ref('')
        const minecraftuuid = ref('')
        const steamid = ref('')

        const submitAppeal = async () => {
            try {
                await axios.post('/api/appeal', {
                    username: username.value,
                    reason: reason.value,
                    minecraftuuid: minecraftuuid.value,
                    steamid: steamid.value
                })
                alert('Appeal submitted successfully')
            } catch (error) {
                console.error('Error submitting appeal:', error)
                alert('An error occurred while submitting the appeal')
            }
        }

        return {
            username,
            reason,
            minecraftuuid,
            steamid,
            submitAppeal
        }
    }
})
</script>

<style>
.appeal {
    max-width: 800px;
    margin: 0 auto;
    text-align: center;
}

.appeal h2 {
    font-size: 2em;
}

.field {
    margin-bottom: 20px;
}

label {
    display: block;
    margin-bottom: 5px;
}

input,
textarea {
    width: 100%;
    padding: 10px;
    font-size: 1em;
    border: 1px solid var(--accent-color);
    border-radius: 5px;
}

button {
    padding: 10px 20px;
    font-size: 1em;
    background: var(--accent-color);
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}
</style>