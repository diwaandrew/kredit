import React, { useState } from 'react'
import { Link } from 'react-router-dom';
import { Menu } from './menu';
// import SidebarMenu from 'react-bootstrap-sidebar-menu';
import { IconContext } from 'react-icons';
import * as FcIcons from "react-icons/fc";
import './Sidebars.css'

const Sidebars = ({ selectData }) => {
    const [sidebar, setsidebar] = useState(true)

    const showSidebar = () => {
        setsidebar(!sidebar)
        selectData(!sidebar)
    }
    const logout = () => {
      localStorage.removeItem('name');
      window.location.href="/"   
    }
    return (
        <>
            <IconContext.Provider value={{ color: 'white' }}>
                <nav className={sidebar ? 'nav-menu active' : 'nav-menu'}>
                    <ul className='nav-menu-items' onClick={showSidebar}>
                        <li className='navbar-toggle'>
                            <Link to='#' className='menu-bars'>
                                <FcIcons.FcMenu />
                            </Link>
                        </li>
                        {Menu.map((item, index) => {
                            return (
                                <li key={index} className={item.cName}>
                                    <Link to={item.path}>
                                        {item.icon}
                                        <span className='span'>{item.title}</span>
                                    </Link>
                                </li>
                            );
                        })}
                    </ul>
                    <div>
                        <li className='nav-text position-absolute bottom-0'>
                            <Link onClick={() => logout()}>
                                <FcIcons.FcImport />
                                <span className='span'>Log Out</span>
                            </Link>
                        </li>
                    </div>
                </nav>
            </IconContext.Provider>
        </>
    )
}

export default Sidebars 