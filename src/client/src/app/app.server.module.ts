import { NgModule } from '@angular/core';
import { ServerModule } from '@angular/platform-server';

import { AppModule } from './app.module';
import { AppComponent } from './app.component';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { CookieInterceptor } from './cookie.interceptor';
import { XhrFactory } from '@angular/common';
//@ts-ignore
import  * as xhr2 from 'xhr2';

export class ServerXhr implements XhrFactory {
  build(): XMLHttpRequest {
    xhr2.prototype._restrictedHeaders.cookie = false;
    return new xhr2.XMLHttpRequest();
  }
}

@NgModule({
  imports: [
    AppModule,
    ServerModule,
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: CookieInterceptor,
      multi: true
    },{
      provide: XhrFactory,
      useClass: ServerXhr
    }
  ],
  bootstrap: [AppComponent],
})
export class AppServerModule {}
