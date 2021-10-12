import { createTheme } from '@mui/material/styles'

export const createAppTheme = (dark = true) => {
    return createTheme({
        palette: {
            mode: dark ? 'dark' : 'light',
        }
    })
}
