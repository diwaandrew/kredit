import React from 'react'
import * as FcIcons from 'react-icons/fc'
export const Menu = [
    {
        title: 'Home',
        path: '/',
        icon: <FcIcons.FcHome />,
        cName: 'nav-text'
    },
    {
        title: 'Checklist',
        path: '/approval',
        icon: <FcIcons.FcInspection />,
        cName: 'nav-text'
    },
    {
        title: 'Drawdown',
        path: '/drawdown',
        icon: <FcIcons.FcDocument />,
        cName: 'nav-text'
    },
    {
        title: 'Change Password',
        path: '/changepassword',
        icon: <FcIcons.FcUnlock />,
        cName: 'nav-text'
    }
]