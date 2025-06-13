import { ref } from 'vue';

export function useFetch(request) {
	const data = ref(null);
	const error = ref(null);
	const loading = ref(false);

	async function execute(customRequest = request) {
		loading.value = true;
		error.value = null;

		try {
			const response = await fetch(customRequest);
			if (!response.ok) {
				throw new Error(`Error: ${response.status} ${response.statusText}`);
			}
			data.value = await response.json();
		} catch (err) {
			// Silly way of handling post requests. Otherwise the backend would have to provide Content-Type headers for correct prasing
			if (err instanceof SyntaxError && err.message.includes('JSON.parse')) {
				error.value = null
			} else {
				error.value = err.message || 'Unknown error';
				console.error(err);
			}
		} finally {
			loading.value = false;
		}
	}

	return {
		data,
		error,
		loading,
		execute,
	};
}
