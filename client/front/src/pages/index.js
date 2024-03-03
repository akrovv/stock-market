import HomeOne from "./HomeOne";
import BuyCrypto from "./BuyCrypto";
import BuyCryptoConfirm from "./BuyCryptoConfirm";
import BuyCryptoDetails from "./BuyCryptoDetails";
import Markets from "./Markets";
import SellSelect from "./SellSelect";
import SellCryptoAmount from "./SellCryptoAmount";
import SellCryptoConfirm from "./SellCryptoConfirm";
import SellCryptoDetails from "./SellCryptoDetails";
import Wallet from "./Wallet";
import UserProfile from "./UserProfile";
import Login from "./Login";
import Register from "./Register";
import Contact from "./Contact";



const routes = [
  { path: '/', component: <HomeOne />},
  { path: '/buy-crypto-select', component: <BuyCrypto />},
  { path: '/buy-crypto-confirm', component: <BuyCryptoConfirm />},
  { path: '/buy-crypto-details', component: <BuyCryptoDetails />},
  { path: '/markets', component: <Markets />},
  { path: '/sell-select', component: <SellSelect />},
  { path: '/sell-crypto-amount', component: <SellCryptoAmount />},
  { path: '/sell-crypto-confirm', component: <SellCryptoConfirm />},
  { path: '/sell-crypto-details', component: <SellCryptoDetails />},
  { path: '/wallet', component: <Wallet />},
  { path: '/user-profile', component: <UserProfile />},
  { path: '/login', component: <Login />},
  { path: '/register', component: <Register />},
  { path: '/contact', component: <Contact />},
]

export default routes;