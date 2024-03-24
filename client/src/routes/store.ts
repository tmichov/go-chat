import { writable } from "svelte/store";

const messageStore = writable('');

const receiveMessage = (message:string) => {
		messageStore.set(message);
}

export default {
		subscribe: messageStore.subscribe,
		receiveMessage
}
