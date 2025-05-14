import { defineConfig, globalIgnores } from 'eslint/config'
import globals from 'globals'
import js from '@eslint/js'
import pluginVue from 'eslint-plugin-vue'
import stylistic from '@stylistic/eslint-plugin'

export default defineConfig([
	...pluginVue.configs['flat/strongly-recommended'],
	js.configs.recommended,
	{
		name: 'app/files-to-lint',
		files: ['**/*.{js,mjs,jsx,vue}'],
		plugins: {
			'@stylistic': stylistic
		},
		rules: {
			"@stylistic/no-multi-spaces": "error",

			"vue/html-indent": ["error", 2],
			"vue/script-indent": ["error", "tab"],
			"vue/singleline-html-element-content-newline": "off",
			"vue/html-closing-bracket-newline" : "off"
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
])
