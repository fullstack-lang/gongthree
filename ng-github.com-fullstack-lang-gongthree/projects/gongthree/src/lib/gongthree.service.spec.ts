import { TestBed } from '@angular/core/testing';

import { GongthreeService } from './gongthree.service';

describe('GongthreeService', () => {
  let service: GongthreeService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongthreeService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
