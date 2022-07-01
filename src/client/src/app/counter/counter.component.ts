import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-counter',
  templateUrl: './counter.component.html',
  styleUrls: ['./counter.component.scss']
})
export class CounterComponent {
 
  @Input() value = 0;
  width = 30;
  height = 30;

  increment(){
    this.value++;
    this.width += (this.value * 4);
    this.height += (this.value * 4);

  }
  decrement(){
    this.width -=this.value * 4;
    this.height -=this.value * 4;
    this.value--;
  }
}
