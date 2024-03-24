<script>
		import {onMount} from 'svelte';
		import Message from './Message.svelte';

		let message = '';
		let messages = [];
		let socket;

		onMount(() => {
				socket = new WebSocket('ws://localhost:8000/ws');

				let connect = cb => {
						console.log('attempting to connect');

						socket.onopen = () => {
								console.log('successfully connected');
						}

						socket.onmessage = msg => {
								console.log(msg);
								cb(msg)
						}

						socket.onclose = () => {
								console.log('socket closed connections');
						}
				}

				connect((msg) => {
						messages = [...messages, msg.data];
				});
		});

		function onSendMessage() {
				if(message.length > 0) {
						socket.send(message);
						message = ''
				}
		}

</script>
<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<input type="text" bind:value={message}>
<button on:click={onSendMessage}>Send Message</button>

<div id="history">
		{#each messages as msg, i}
				<Message message={msg}	direction={i%2 == 0 ? 'left' : 'right'}/>
		{/each}
</div>

<style>
		#history {
				min-height: 300px;
				min-width: 300px;
				max-height: 300px;
				overflow-y: auto;
		}
</style>
