import { defineConfig, globalIgnores } from 'eslint/config'
import globals from 'globals'
import js from '@eslint/js'
import pluginVue from 'eslint-plugin-vue'

export default defineConfig([
	{
		name: 'app/files-to-lint',
		files: ['**/*.{js,mjs,jsx,vue}'],
		rules: {
			// Technically eslint formatting is deprecated but it's good enough for this
			// The alternative is to setup prettier but I don't like it
			"vue/html-indent": ["error", 2],
			"vue/script-indent": ["error", "tab"],
			"no-multi-spaces": "error",
			"no-trailing-spaces": "error"

		}
	},

	globalIgnores(['**/dist/**', '**/dist-ssr/**', '**/coverage/**']),

	{
		languageOptions: {
			globals: {
				...globals.browser,
			},
		},
	},

	js.configs.recommended,
	...pluginVue.configs['flat/essential'],
])
