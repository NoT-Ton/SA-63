import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Output';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ใบบันทึกการขาย' };
 
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={` ${profile.givenName || 'to Backstage'}`}
       subtitle="บอกข้อมูลสินค้าที่ถูกขาย"
        
           
         
     ></Header>
     <Content>
       <ContentHeader title="">
         <Link component={RouterLink} to="/input">
           <Button variant="contained" color="primary">
             Add Data
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;