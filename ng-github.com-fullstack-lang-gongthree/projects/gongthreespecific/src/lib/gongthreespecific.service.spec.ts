import { TestBed } from '@angular/core/testing';

import { GongthreespecificService } from './gongthreespecific.service';

describe('GongthreespecificService', () => {
  let service: GongthreespecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongthreespecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
