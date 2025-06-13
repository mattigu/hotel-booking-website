import { ref } from 'vue';

export function useFetch(request, options = { immediate: true }) {
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
			error.value = err.message || 'Unknown error';
			console.error(err);
		} finally {
			loading.value = false;
		}
	}

	if (options.immediate) {
		execute();
	}

	return {
		data,
		error,
		loading,
		execute,
	};
}
