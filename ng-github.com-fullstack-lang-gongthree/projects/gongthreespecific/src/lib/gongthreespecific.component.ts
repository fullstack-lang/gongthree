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
    this.camera.position.z = 5;

    // Create renderer
    this.renderer = new THREE.WebGLRenderer({
      canvas: this.canvasRef.nativeElement,
      antialias: true
    });
    this.renderer.setSize(this.canvasRef.nativeElement.clientWidth, this.canvasRef.nativeElement.clientHeight);
    this.renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
  }

  private createObjects() {
    // Create a colorful cube
    const geometry = new THREE.BoxGeometry(1, 1, 1);
    const materials = [
      new THREE.MeshStandardMaterial({ color: 0xff0000 }), // red
      new THREE.MeshStandardMaterial({ color: 0x00ff00 }), // green
      new THREE.MeshStandardMaterial({ color: 0x0000ff }), // blue
      new THREE.MeshStandardMaterial({ color: 0xffff00 }), // yellow
      new THREE.MeshStandardMaterial({ color: 0xff00ff }), // magenta
      new THREE.MeshStandardMaterial({ color: 0x00ffff })  // cyan
    ];
    const cube = new THREE.Mesh(geometry, materials);
    this.scene.add(cube);

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

    // Extrude settings
    const extrudeSettings: THREE.ExtrudeGeometryOptions = {
      depth: height,
      bevelEnabled: false,
      steps: 1,
    };

    // Create the geometry
    const hollowCylinderGeometry = new THREE.ExtrudeGeometry(shape, extrudeSettings);

    // Rotate the geometry to align it along the Y-axis
    hollowCylinderGeometry.rotateX(Math.PI / 2);

    // Create the mesh
    const hollowCylinderMaterial = new THREE.MeshStandardMaterial({ color: 0x00ff00 });
    const hollowCylinder = new THREE.Mesh(hollowCylinderGeometry, hollowCylinderMaterial);
    hollowCylinder.position.set(1.5, 0, 0); // Position cylinder slightly to the right
    hollowCylinder.position.y -= height / 2; // Adjust position since extrusion starts from z=0

    // Add to the scene
    this.scene.add(hollowCylinder);
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
}
