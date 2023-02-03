import { TestBed } from '@angular/core/testing';

import { ItemDeleterService } from './item-deleter.service';

describe('ItemDeleterService', () => {
  let service: ItemDeleterService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ItemDeleterService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
