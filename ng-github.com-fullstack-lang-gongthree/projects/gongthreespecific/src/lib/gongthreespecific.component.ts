import { Component, ElementRef, OnInit, ViewChild, OnDestroy } from '@angular/core';

import * as THREE from 'three';

import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';

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
    this.camera.position.z = 15;

    // Create renderer
    this.renderer = new THREE.WebGLRenderer({
      canvas: this.canvasRef.nativeElement,
      antialias: true
    });
    this.renderer.setSize(this.canvasRef.nativeElement.clientWidth, this.canvasRef.nativeElement.clientHeight);
    this.renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
  }

  private createObjects() {

    // Create a hollow cylinder
    const outerRadius = 0.5;
    const wallThickness = 0.05;
    const innerRadius = outerRadius - wallThickness;
    const height = 2;

    // Define the outer shape (circle)
    const shape = new THREE.Shape();
    shape.absarc(0, 0, outerRadius, 0, Math.PI * 2, false);

    // Define the hole (inner circle)
    const holePath = new THREE.Path();
    holePath.absarc(0, 0, innerRadius, 0, Math.PI * 2, true);
    shape.holes.push(holePath);

    // Extrude settings with smooth rendering
    const extrudeSettings: THREE.ExtrudeGeometryOptions = {
      depth: height,
      bevelEnabled: false,
      steps: 1,
      curveSegments: 64 // Increase the number of segments for smoothness
    };

    // Create the geometry
    const hollowCylinderGeometry = new THREE.ExtrudeGeometry(shape, extrudeSettings);

    // Rotate the geometry to align it along the Y-axis
    hollowCylinderGeometry.rotateX(Math.PI / 2);

    // Create the mesh
    const hollowCylinderMaterial = new THREE.MeshStandardMaterial({ color: 0x00ff00 });
    const hollowCylinder = new THREE.Mesh(hollowCylinderGeometry, hollowCylinderMaterial);
    hollowCylinder.position.set(1.5, 2, 0); // Position cylinder slightly to the right
    hollowCylinder.position.y -= height / 2; // Adjust position since extrusion starts from z=0

    // Add to the scene
    this.scene.add(hollowCylinder);

    this.createRandomExtrudedShape()

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

  private createRandomExtrudedShape() {
    // Generate random points on a circle
    const radius = 10;
    const variation = 0;
    const randomRadius = () => radius + (Math.random() * variation * 2 - variation);

    const points: THREE.Vector2[] = [];
    const nbPoints = 3; // Number of random points
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
    const material = new THREE.MeshStandardMaterial({ color: 0xff00ff });
    const extrudedMesh = new THREE.Mesh(extrudedGeometry, material);

    // Add the extruded shape to the scene
    this.scene.add(extrudedMesh);
  }



}
