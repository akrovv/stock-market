import React , {useState, useEffect} from 'react';

import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import PageTitle from '../components/pagetitle';
import {Link} from 'react-router-dom';
import img from '../assets/images/avt/avt.png'

UserProfile.propTypes = {
};

async function loginAndFetchData(setUserInfo) {
    try {
        const response = await fetch("/api/user/profile", {
            method: "GET",
        });

        if (!response.ok) {
            throw new Error("Network response was not ok");
        }

        const data = await response.json();
        console.log(data, response.status);

        setUserInfo({
            nickname: data.nickname,
            email: data.email,
            firstname: data.firstname,
            lastname: data.lastname
        });
    } catch (error) {
        console.error(error);
    }
}


function UserProfile(props) {
    const [dataCoinTab] = useState([
        {
            id: 1,
            title: 'User Profile',
            icon: 'fa-user'
        },
        {
            id: 2,
            title: 'Referrals',
            icon: 'fa-share-nodes'
        },
        {
            id: 5,
            title: 'Change password',
            icon: 'fa-lock'
        },
    ]);

    const [userInfo, setUserInfo] = useState({
        nickname: "",
        email: "",
        firstname: "",
        lastname: ""
    });

    useEffect(() => {
        loginAndFetchData(setUserInfo);
    }, []);
    

    const { nickname, email, firstname, lastname } = userInfo;
    
    return (
        <div>
            <PageTitle heading='User Profile' title='User' />


            <section className="user-profile flat-tabs">
            <div className="container">
                <div className="row">
                <Tabs>
                    
                    <TabList>
                        <div className="user-info center">
                            <div className="avt">
                                <input
                                type="file"
                                className="custom-file-input"
                                id="imgInp"
                                required
                                />
                                <img id="blah" src={img} alt="no file" />
                            </div>
                            <h6 className="name" id="nickname">{nickname}</h6>
                            <p id="email">{email}</p>
                        </div>
                        {
                            dataCoinTab.map(idx => (
                                <Tab key={idx.id}><h6 className="fs-16">
                                <i className={`fa ${idx.icon}`}></i>
                                {idx.title}
                                </h6></Tab>
                            ))
                        }

                    </TabList>

                    <TabPanel>
                        <div className="content-inner profile">
                            <form action="#">
                            <h4>User Profile</h4>
                            <h6>Infomation</h6>

                            <div className="form-group d-flex s1">
                                FirstName: <input type="text" className="form-control" id="firstname" value={firstname} required/>
                                LastName: <input
                                type="text"
                                className="form-control"
                                id="lastname"
                                value={lastname}
                                required
                                />
                            </div>

                            <button type="submit" className="btn-action">
                                Update Profile
                            </button>
                            </form>
                        </div>
                    </TabPanel>

                    <TabPanel>
                        <div className="content-inner referrals">
                            <h6>Total rewards</h6>
                            <h4>$1,056.00 <span>USD</span></h4>
                            <p>
                            You're earning 20% of the trading fees your referrals pay.
                            Learn more
                            </p>
                            <div className="main">
                            <h6>Invite friends to earn 20%</h6>

                            <div className="refe">
                                <div>
                                <p>Referral link</p>
                                <input
                                    className="form-control"
                                    type="text"
                                    value="https://accounts.rockie.com/login"
                                />
                                </div>
                                <div>
                                <p>Referral code</p>
                                <input
                                    className="form-control"
                                    type="text"
                                    value="N84CRDKK"
                                />
                                <span className="btn-action">Copied</span>
                                </div>
                            </div>
                            </div>

                            <Link to="/wallet" className="btn-action">My Wallet</Link>
                        </div>
                    </TabPanel>
                    <TabPanel>
                        <div className="content-inner profile change-pass">
                            <h4>Change Password</h4>
                            <h6>New Passworld</h6>
                            <form action="#">
                            <div className="form-group">
                                <div>
                                <label>Old Passworld<span>*</span>:</label>
                                <input
                                    type="text"
                                    className="form-control"
                                    value="123456789"
                                />
                                </div>
                            </div>
                            <div className="form-group">
                                <div>
                                <label>New Passworld<span>*</span>:</label>
                                <input
                                    type="password"
                                    className="form-control"
                                    placeholder="New Passworld"
                                />
                                </div>
                                <div>
                                <label>Confirm Passworld<span>*</span>:</label>
                                <input
                                    type="password"
                                    className="form-control"
                                    placeholder="Confirm Passworld"
                                />
                                </div>
                            </div>
                            </form>
                            <button type="submit" className="btn-action">
                            Change Passworld
                            </button>
                        </div>
                    </TabPanel>
                    

                </Tabs> 
                </div>
            </div>
            </section>
            
        </div>
    );
}

export default UserProfile;