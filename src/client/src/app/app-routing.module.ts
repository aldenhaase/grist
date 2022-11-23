import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './auth/login/login.component';
import { ListComponent } from './home/list/list.component';
import { RegisterComponent } from './auth/register/register.component';
import { AuthGuard } from './guards/auth.guard';

const routes: Routes = [
  {path: '', 
  component: ListComponent,
  canActivate: [AuthGuard]},
  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegisterComponent},
]

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule { }
