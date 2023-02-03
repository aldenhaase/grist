import { TestBed } from '@angular/core/testing';

import { ListSetterService } from './list-setter.service';

describe('ListSetterService', () => {
  let service: ListSetterService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ListSetterService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
