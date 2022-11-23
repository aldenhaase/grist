import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AuthModule } from './auth/auth.module';
import { HomeModule } from './home/home.module';
import {HttpClientModule } from '@angular/common/http';
import { AuthStatusService } from './services/auth-status.service';
import { AuthGuard } from './guards/auth.guard';

@NgModule({
  declarations: [
    AppComponent
    ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    AuthModule,
    HomeModule,
    HttpClientModule
  ],
  providers: [AuthStatusService,AuthGuard],
  bootstrap: [AppComponent]
})
export class AppModule { }
