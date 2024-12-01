import { Component, ElementRef, OnInit, ViewChild, OnDestroy } from '@angular/core';

import * as THREE from 'three';

import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';
import { BezierSegment } from './bezier-segment';

@Component({
  selector: 'lib-gongthreespecific',
  standalone: true,
  imports: [],
  templateUrl: './gongthreespecific.component.html',
  styleUrls: ['./gongthreespecific.component.css'],
})
export class GongthreespecificComponent {

  @ViewChild('threeCanvas', { static: true }) canvasRef!: ElementRef<HTMLCanvasElement>;

  private scene!: THREE.Scene;
  private camera!: THREE.PerspectiveCamera;
  private renderer!: THREE.WebGLRenderer;
  private controls!: OrbitControls;
  private animationFrameId?: number;

  ngOnInit() {
    this.initScene();
    this.createObjects();
    this.setupLighting();
    this.setupControls();
    this.animate();
  }

  private initScene() {
    // Create scene
    this.scene = new THREE.Scene();
    this.scene.background = new THREE.Color(0xf0f0f0);

    // Create camera
    const aspectRatio = this.getAspectRatio();
    this.camera = new THREE.PerspectiveCamera(75, aspectRatio, 0.1, 1000);
    this.camera.position.z = 25;

    // Create renderer
    this.renderer = new THREE.WebGLRenderer({
      canvas: this.canvasRef.nativeElement,
      antialias: true
    });
    this.renderer.setSize(this.canvasRef.nativeElement.clientWidth, this.canvasRef.nativeElement.clientHeight);
    this.renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
  }

  private createObjects() {

    // ['#D1C5B4', '#8FA382', '#536C87'],

    // this.createRandomExtrudedShape(15, 0.5, 5, 0, "#D1C5B4")

    // this.createRandomExtrudedShape(10, 0.5, 5, 1, "#8FA382")

    // this.createRandomExtrudedShape(8, 0.5, 8, 2, "#536C87")

    // Define multiple bezier segments
    const bezierSegments: BezierSegment[] = [
      new BezierSegment(
        new THREE.Vector2(-5, 4),
        new THREE.Vector2(-2, 8),
        new THREE.Vector2(2, 8),
        new THREE.Vector2(5, 4)
      ),
      new BezierSegment(
        new THREE.Vector2(5, 4),
        new THREE.Vector2(18, 0),
        new THREE.Vector2(8, -4),
        new THREE.Vector2(5, -8)
      ),
      new BezierSegment(
        new THREE.Vector2(5, -8),
        new THREE.Vector2(2, -12),
        new THREE.Vector2(-2, -12),
        new THREE.Vector2(-5, -8)
      ),
      new BezierSegment(
        new THREE.Vector2(-5, -8),
        new THREE.Vector2(-8, -4),
        new THREE.Vector2(-8, 0),
        new THREE.Vector2(-5, 4)
      ),
    ];

    // Call the method with the array of segments
    this.createBezierShape(bezierSegments);



    // Add axes helper
    const axesHelper = new THREE.AxesHelper(2); // Length of the axes lines
    this.scene.add(axesHelper);
  }

  private setupLighting() {
    // Add multiple lights to create depth and shadows
    const ambientLight = new THREE.AmbientLight(0xffffff, 0.5);
    this.scene.add(ambientLight);

    const directionalLight = new THREE.DirectionalLight(0xffffff, 0.8);
    directionalLight.position.set(5, 5, 5);
    this.scene.add(directionalLight);
  }

  private setupControls() {
    // Add orbit controls to allow interactive camera movement
    this.controls = new OrbitControls(this.camera, this.renderer.domElement);
    this.controls.enableDamping = true;
  }

  private animate() {
    // Render loop with animation
    const renderLoop = () => {
      // Update controls
      this.controls.update();

      // Render the scene
      this.renderer.render(this.scene, this.camera);

      // Request next frame
      this.animationFrameId = requestAnimationFrame(renderLoop);
    };

    renderLoop();
  }

  private getAspectRatio(): number {
    return this.canvasRef.nativeElement.clientWidth / this.canvasRef.nativeElement.clientHeight;
  }

  ngOnDestroy() {
    // Clean up resources
    if (this.animationFrameId) {
      cancelAnimationFrame(this.animationFrameId);
    }
    this.controls.dispose();
    this.renderer.dispose();
  }

  private createRandomExtrudedShape(
    refRadius: number,
    variation: number,
    nbPoints: number,
    zOrig: number,
    color: string) {
    // Generate random points on a circle
    const radius = refRadius;
    const randomRadius = () => radius + (Math.random() * variation * 2 - variation);

    const points: THREE.Vector2[] = [];
    for (let i = 0; i < nbPoints; i++) {
      const angle = (i * Math.PI * 2) / nbPoints;
      const r = randomRadius();
      const x = Math.cos(angle) * r;
      const y = Math.sin(angle) * r;
      points.push(new THREE.Vector2(x, y));
    }

    // Convert points into Bézier segments
    const shape = new THREE.Shape();
    const controlPointDistance = radius * 0.5; // Distance of control points from the curve points

    // Start with the first point
    shape.moveTo(points[0].x, points[0].y);

    for (let i = 0; i < points.length; i++) {
      const p0 = points[i];
      const p1 = points[(i + 1) % points.length]; // Next point (wrapping around)

      // Vector from the center to p0
      const tangent0 = new THREE.Vector2(-p0.y, p0.x).normalize(); // Rotate 90 degrees
      const tangent1 = new THREE.Vector2(-p1.y, p1.x).normalize(); // Rotate 90 degrees

      // Control points
      const cp0 = new THREE.Vector2(
        p0.x + tangent0.x * controlPointDistance,
        p0.y + tangent0.y * controlPointDistance
      );
      const cp1 = new THREE.Vector2(
        p1.x - tangent1.x * controlPointDistance,
        p1.y - tangent1.y * controlPointDistance
      );

      // Cubic Bézier curve
      shape.bezierCurveTo(cp0.x, cp0.y, cp1.x, cp1.y, p1.x, p1.y);
    }

    shape.closePath();

    // Extrude the shape along the Z-axis
    const extrudeSettings: THREE.ExtrudeGeometryOptions = {
      depth: -0.5,
      bevelEnabled: false,
      steps: 1,
    };

    const extrudedGeometry = new THREE.ExtrudeGeometry(shape, extrudeSettings);
    const material = new THREE.MeshStandardMaterial({ color: color });
    const extrudedMesh = new THREE.Mesh(extrudedGeometry, material);

    // Set the position of the shape (including the Z parameter)
    extrudedMesh.position.set(0, 0, zOrig)

    // Add the extruded shape to the scene
    this.scene.add(extrudedMesh);
  }

  private createBezierShape(bezierSegments: BezierSegment[]) {
    // Check if the array is not empty
    if (bezierSegments.length === 0) {
      console.warn('No bezier segments provided.');
      return;
    }

    // Create a shape
    const shape = new THREE.Shape();

    // Move to the starting point of the first segment
    const firstSegment = bezierSegments[0];
    shape.moveTo(firstSegment.Start.x, firstSegment.Start.y);

    // Create spheres to mark important points
    const createPointMarker = (position: THREE.Vector2, color: number) => {
      const sphereGeometry = new THREE.SphereGeometry(0.2, 32, 32);
      const sphereMaterial = new THREE.MeshStandardMaterial({ color: color });
      const sphere = new THREE.Mesh(sphereGeometry, sphereMaterial);
      sphere.position.set(position.x, position.y, 0.1); // Slightly above the shape
      this.scene.add(sphere);
    };

    // Mark start point (red)
    createPointMarker(firstSegment.Start, 0xff0000);

    // Loop through each bezier segment and add it to the shape
    for (let i = 0; i < bezierSegments.length; i++) {
      const segment = bezierSegments[i];

      // If this is not the first segment, ensure continuity
      if (i > 0) {
        // Optionally, move to the start point of the segment if it's not connected
        // shape.lineTo(segment.Start.x, segment.Start.y);
      }

      // Add bezier curve to the shape
      shape.bezierCurveTo(
        segment.ControlPointStart.x, segment.ControlPointStart.y,
        segment.ControlPointEnd.x, segment.ControlPointEnd.y,
        segment.End.x, segment.End.y
      );

      // Mark control points (blue)
      createPointMarker(segment.ControlPointStart, 0x0000ff);
      createPointMarker(segment.ControlPointEnd, 0x0000ff);

      createPointMarker(segment.End, 0x00ff00);

    }

    // Optionally close the shape
    // shape.closePath();

    // Create geometry and material for the shape
    const geometry = new THREE.ShapeGeometry(shape);
    const material = new THREE.MeshStandardMaterial({
      color: 0x8FA382,
      side: THREE.DoubleSide,
      opacity: 1,
      transparent: false
    });
    const mesh = new THREE.Mesh(geometry, material);

    // Adjust position to be more visible
    mesh.position.z = 0;
    mesh.rotation.x = 0;

    // Add to the scene
    this.scene.add(mesh);

    // Optional: Add wireframe to help visualize shape
    const wireframe = new THREE.LineSegments(
      new THREE.EdgesGeometry(geometry),
      new THREE.LineBasicMaterial({ color: 0x000000, linewidth: 2 })
    );
    this.scene.add(wireframe);
  }





}
