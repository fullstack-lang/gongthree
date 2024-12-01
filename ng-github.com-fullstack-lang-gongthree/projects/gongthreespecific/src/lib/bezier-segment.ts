import * as THREE from 'three';

export class BezierSegment {
    Start: THREE.Vector2;
    ControlPointStart: THREE.Vector2;
    ControlPointEnd: THREE.Vector2;
    End: THREE.Vector2;

    constructor(
        Start: THREE.Vector2,
        ControlPointStart: THREE.Vector2,
        ControlPointEnd: THREE.Vector2,
        End: THREE.Vector2
    ) {
        this.Start = Start;
        this.ControlPointStart = ControlPointStart;
        this.ControlPointEnd = ControlPointEnd;
        this.End = End;
    }
}

// Usage:
const bezierSegment = new BezierSegment(
    new THREE.Vector2(-5, 4),
    new THREE.Vector2(-2, 8),
    new THREE.Vector2(2, 8),
    new THREE.Vector2(5, 4)
);
