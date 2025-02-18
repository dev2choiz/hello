import { styled, Theme, CSSObject, ThemeProvider } from '@mui/material/styles'
import Box from '@mui/material/Box'
import MuiDrawer from '@mui/material/Drawer'
import MuiAppBar, { AppBarProps as MuiAppBarProps } from '@mui/material/AppBar'
import Toolbar from '@mui/material/Toolbar'
import List from '@mui/material/List'
import CssBaseline from '@mui/material/CssBaseline'
import Typography from '@mui/material/Typography'
import Divider from '@mui/material/Divider'
import IconButton from '@mui/material/IconButton'
import MenuIcon from '@mui/icons-material/Menu'
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft'
import ChevronRightIcon from '@mui/icons-material/ChevronRight'
import AcUnitIcon from '@mui/icons-material/AcUnit'
import HealthAndSafetyIcon from '@mui/icons-material/HealthAndSafety'
import AccessibilityIcon from '@mui/icons-material/Accessibility'
import AccessTimeIcon from '@mui/icons-material/AccessTime'
import AccountBalanceIcon from '@mui/icons-material/AccountBalance'
import BrightnessHighIcon from '@mui/icons-material/BrightnessHigh'
import BrightnessLowIcon from '@mui/icons-material/BrightnessLow'
import ListItem from '@mui/material/ListItem'
import ListItemIcon from '@mui/material/ListItemIcon'
import ListItemText from '@mui/material/ListItemText'
import InboxIcon from '@mui/icons-material/MoveToInbox'
import MailIcon from '@mui/icons-material/Mail'
import { PropsWithChildren, useEffect, useState } from 'react'
import faker from 'faker'
import { createAppTheme } from '@/styles/theme'
import Link from 'next/link'

const drawerWidth = 240

const openedMixin = (theme: Theme): CSSObject => ({
    width: drawerWidth,
    transition: theme.transitions.create('width', {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
    }),
    overflowX: 'hidden',
})

const closedMixin = (theme: Theme): CSSObject => ({
    transition: theme.transitions.create('width', {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
    }),
    overflowX: 'hidden',
    width: `calc(${theme.spacing(7)} + 1px)`,
    [theme.breakpoints.up('sm')]: {
        width: `calc(${theme.spacing(9)} + 1px)`,
    },
})

const DrawerHeader = styled('div')(({ theme }) => ({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: theme.spacing(0, 1),
    // necessary for content to be below app bar
    ...theme.mixins.toolbar,
}))

interface AppBarProps extends MuiAppBarProps {
    open?: boolean;
}

const AppBar = styled(MuiAppBar, {
    shouldForwardProp: (prop) => prop !== 'open',
})<AppBarProps>(({ theme, open }) => ({
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
    }),
    ...(open && {
        marginLeft: drawerWidth,
        width: `calc(100% - ${drawerWidth}px)`,
        transition: theme.transitions.create(['width', 'margin'], {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.enteringScreen,
        }),
    }),
}))

const Drawer = styled(MuiDrawer, { shouldForwardProp: (prop) => prop !== 'open' })(
    ({ theme, open }) => ({
        width: drawerWidth,
        flexShrink: 0,
        whiteSpace: 'nowrap',
        boxSizing: 'border-box',
        ...(open && {
            ...openedMixin(theme),
            '& .MuiDrawer-paper': openedMixin(theme),
        }),
        ...(!open && {
            ...closedMixin(theme),
            '& .MuiDrawer-paper': closedMixin(theme),
        }),
    }),
)

const Layout = ({ children }: PropsWithChildren< Record<string, any> >) => {
    const [open, setOpen] = useState(true)
    const [darkMode, setDarkMode] = useState(true)

    const handleDrawerOpen = () => {
        setOpen(true)
    }

    const [randomName, setRandomName] = useState<string>('rand')
    const handleDrawerClose = () => {
        setOpen(false)
    }

    const theme = createAppTheme(darkMode)

    useEffect(() => {
        setRandomName(faker.name.lastName())
        const interval = setInterval(() => { setRandomName(faker.name.lastName()) }, 5000)
        return () => { clearInterval(interval) }
    }, [])

    const links = [
        ['ISR - health', '/health', '/health', <HealthAndSafetyIcon key={0}/>],
        ['SSR', '/', '/', <InboxIcon key={1}/>],
        ['SSR - with params', '/?name=John', '/?name=John', <MailIcon key={2}/>],
        ['SSG', '/unary-static', '/unary-static', <AcUnitIcon key={3}/>],
        ['ISG - with params', `/unary-static/${randomName}`, `/unary-static/${randomName}`, <AccessibilityIcon key={4}/>],
        ['SPA - server stream', '/server-stream', '/server-stream', <AccessTimeIcon key={5}/>],
        ['SPA - server stream with param', '/server-stream/5', '/server-stream/5', <AccountBalanceIcon key={6}/>],
    ]

    return (
        <ThemeProvider theme={theme}>
            {/*CssBaseline kickstart an elegant, consistent, and simple baseline to build upon.*/}
            <CssBaseline />
            <Box sx={{ display: 'flex' }}>
                <CssBaseline />
                <AppBar position="fixed" open={open}>
                    <Toolbar>
                        <IconButton
                            color="inherit"
                            aria-label="open drawer"
                            onClick={handleDrawerOpen}
                            edge="start"
                            sx={{
                                marginRight: '36px',
                                ...(open && { display: 'none' }),
                            }}
                        >
                            <MenuIcon />
                        </IconButton>
                        <Box sx={{ flexGrow: 1 }} />
                        <Typography variant="h6" noWrap component="div">
                            HelloFront
                        </Typography>
                        <Box sx={{ flexGrow: 1 }} />
                        <IconButton
                            color="inherit"
                            aria-label="open drawer"
                            onClick={() => { setDarkMode(!darkMode) }}
                            edge="end"
                        >
                            { darkMode ? <BrightnessLowIcon /> : <BrightnessHighIcon /> }
                        </IconButton>
                    </Toolbar>
                </AppBar>
                <Drawer variant="permanent" open={open}>
                    <DrawerHeader>
                        <IconButton onClick={handleDrawerClose}>
                            {theme.direction === 'rtl' ? <ChevronRightIcon /> : <ChevronLeftIcon />}
                        </IconButton>
                    </DrawerHeader>
                    <Divider />
                    <List>
                        {links.map((l, i) => (
                            <Link
                                href={l[1] as string}
                                as={l[2] as string}
                                key={i}
                                prefetch={false}
                                passHref
                            >
                                <ListItem
                                    component={'a'}
                                    button={true}
                                >
                                    <ListItemIcon>{l[3]}</ListItemIcon>
                                    <ListItemText primary={l[0]} secondary={l[1]} />
                                </ListItem>
                            </Link>
                        ))}
                    </List>
                </Drawer>
                <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
                    <DrawerHeader />
                    { children }
                </Box>
            </Box>
        </ThemeProvider>
    )
}

export default Layout
