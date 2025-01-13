<template>
    <section class="bans">
        <h2>Recent Bans</h2>
        <div class="bans-wrapper">
        </div>
    </section>
</template>

<script lang="ts">
import { onMounted } from 'vue'
import axios from 'axios'

onMounted(async () => {
  try {
    const response = await axios.get('/api/bans/recent')
    const data = response.data
    if (data.bans) {
        const bansWrapper = document.querySelector('.bans-wrapper')
        if (bansWrapper) {
            data.bans.forEach((ban: any) => {
            const banElement = document.createElement('div')
            banElement.classList.add('ban')
            banElement.innerHTML = `
                <h3>${ban.identifier}</h3>
                <p>${ban.reason}</p>
                <span>${new Date(ban.date_banned).toLocaleString()}</span>
                <span>${ban.game}</span>
                <span>${ban.steam_id}</span>
                <span>${ban.minecraft_uuid}</span>
            `
            bansWrapper.appendChild(banElement)
            })
        }
    }
  } catch (error) {
    console.error('Error fetching recent bans:', error)
  }
})

</script>

<style>
.bans {
    max-width: 800px;
    margin: 0 auto;
    text-align: center;
}

.bans h2 {
    font-size: 2em;
}

.bans-wrapper {
    display: grid;
    gap: 20px;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
}

.ban {
    background: var(--background-color);
    padding: 20px;
    border: 1px solid var(--accent-color);
    border-radius: 5px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    text-align: center;
}

.ban h3 {
    font-size: 1.5em;
}

.ban p {
    font-size: 1.2em;
}

.ban span {
    display: block;
    margin-top: 10px;
}

.ban hr {
    margin: 10px 0;
}

code {
    font-size: 0.7em;
}

@media (max-width: 768px) {
    .bans-wrapper {
        grid-template-columns: 1fr;
    }
}

@media (min-width: 1024px) {
    .bans-wrapper {
        gap: 40px;
    }
}


</style>