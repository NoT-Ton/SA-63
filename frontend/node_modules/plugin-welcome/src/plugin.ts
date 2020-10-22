import { createPlugin } from '@backstage/core';
import CreateUser from './components/Inputs';
import LogIn from './components/Login';
import WelcomePage from './components/WelcomePage';

 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    
    router.registerRoute('/output', WelcomePage); 
    router.registerRoute('/input', CreateUser);
    router.registerRoute('/', LogIn);
  },
});
